package comms

import (
        "github.com/gorilla/websocket"
        "fmt"
        "net/http"
        "encoding/json"
        "time"
        "log"
        "strconv"
        "os/signal"
        "os"
        "syscall"
        "github.com/eriklupander/ocpvizz/ocpclient"
        "github.com/spf13/viper"
	"strings"
)

type IEventServer interface {
        AddEventToSendQueue(data []byte)
        Close()
}

func NewEventServer() IEventServer {
        es := &EventServer{}
        es.initEventServer()
        return es
}

type EventServer struct {
        upgrader           websocket.Upgrader
        // Create unbuffered channel
        eventQueue         chan []byte
        // Web Socket connection registry (in case we have > 1 dashboards driven by this backend)
        connectionRegistry []*websocket.Conn
}

func (server *EventServer) initEventServer() {
        server.upgrader = websocket.Upgrader{} // use default options
        server.connectionRegistry = make([]*websocket.Conn, 0, 10)
        server.eventQueue = make(chan []byte, 100)
        go server.initializeEventSystem()
}

func (server *EventServer) AddEventToSendQueue(data []byte) {
        server.eventQueue <- data
}

func (server *EventServer) initializeEventSystem() {

        fmt.Println("Starting WebSocket server at port 6969")

        http.HandleFunc("/start", server.registerChannel)
        http.HandleFunc("/nodes", server.getNodes)
        http.HandleFunc("/services", server.getServices)
        http.HandleFunc("/tasks", server.getTasks)
        http.HandleFunc("/containers", server.getContainers)
        http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
                http.ServeFile(w, r, "static/" + r.URL.Path[1:])
        })

        go server.startEventSender()

        handleSigterm(func() {
                server.Close()
        })
        fmt.Println("Registered sigterm handler")

        fmt.Println("Starting WebSocket server")
        err := http.ListenAndServe(":6969", nil)
        if err != nil {
                panic("ListenAndServe: " + err.Error())
        }
}

func (server *EventServer) Close() {
        for _, wsConn := range server.connectionRegistry {
                adr := wsConn.RemoteAddr().String()
                wsConn.Close()
                fmt.Println("Gracefully shut down websocket connection to " + adr)
        }
        for index, _ := range server.connectionRegistry {
                server.removeConnection(index)
                fmt.Printf("Removed connection object at index %v", index)
        }

}

func (server *EventServer) getNodes(w http.ResponseWriter, r *http.Request) {
        //nodes, err := server.Client.ListNodes(docker.ListNodesOptions{})
        //if err != nil {
        //        panic(err)
        //}
        nodes := make([]Project, 0)

        project := &Project {
			ID: viper.GetString("project.name"),
			Description: Desc {Hostname: "hostname"},
			Status: Status {State: "Running"},
		}
        nodes = append(nodes, *project)
        json, _ := json.Marshal(&nodes)
        writeResponse(w, json)
}

type Desc struct {
	Hostname string
}

type Status struct {
	State string
}

type Project struct {
	ID string
	Description Desc
	Status Status
}

type Service struct {
	ID string
	Spec ServiceSpec
}

type ServiceSpec struct {
	Name string
}

// {"id": item.ID, "name": item.Spec.Name}
func (server *EventServer) getServices(w http.ResponseWriter, r *http.Request) {
        services, err := ocpclient.GetServices(viper.GetString("server.url"), viper.GetString("project.name")) //client.ListTasks(docker.ListTasksOptions{Filters: filters})
        //server.Client.ListServices(docker.ListServicesOptions{})
        if err != nil {
                panic(err)
        }
        results := make([]Service, 0)
        for _, service := range services {
        	spec := ServiceSpec{Name: service.Metadata.Name,}
        	serv  := Service{ID: service.Metadata.Name, Spec: spec,}
        	results = append(results, serv)
		}
        json, _ := json.Marshal(&results)
        writeResponse(w, json)
}

/*
if (item.DesiredState !== 'running') {
	return;
}
tasks.push({
	"id": item.ID,
	"image": item.Spec.ContainerSpec.Image,
	"name": item.Spec.ContainerSpec.Image.slice(0, item.Spec.ContainerSpec.Image.indexOf('@')) + "." + item.Slot,
	"serviceId": item.ServiceID,
	"serviceName": item.Spec.ContainerSpec.Image,
	"nodeId": item.NodeID,
	"status": item.Status.State
})
 */

 type Spec struct {
 	ContainerSpec PContainerSpec
 }

 type PContainerSpec struct {
 	Image string
 }

 type Pod struct {
 	ID string
 	DesiredState string
 	Slot string
 	Spec Spec
 	ServiceID string
 	NodeID string
 	Status Status
 }

func (server *EventServer) getTasks(w http.ResponseWriter, r *http.Request) {
        tasks, err := ocpclient.GetPods(viper.GetString("server.url"), viper.GetString("project.name")) //client.ListTasks(docker.ListTasksOptions{Filters: filters})
        //server.Client.ListTasks(docker.ListTasksOptions{})
        if err != nil {
                panic(err)
        }
        pods := make([]Pod, 0)
        for _, task := range tasks {

        	pspec := PContainerSpec{Image: "@" + task.Status.ContainerStatuses[0].Image,}
        	spec := Spec{ContainerSpec: pspec,}
        	status := Status{State: strings.ToLower(task.Status.Phase),}
        	pod := Pod{ID: task.Metadata.UID, DesiredState: strings.ToLower(task.Status.Phase), NodeID: viper.GetString("project.name"), ServiceID: task.Metadata.Labels.Deploymentconfig, Slot: task.Metadata.Name, Status: status, Spec: spec}
			pods = append(pods, pod)
		}
        json, _ := json.Marshal(&pods)
        writeResponse(w, json)
}

func (server *EventServer) getContainers(w http.ResponseWriter, r *http.Request) {
        containers, err := ocpclient.GetPods(viper.GetString("server.url"), viper.GetString("project.name")) //client.ListTasks(docker.ListTasksOptions{Filters: filters})
        //server.Client.ListContainers(docker.ListContainersOptions{All: false})
        if err != nil {
                panic(err)
        }
        json, _ := json.Marshal(&containers)
        writeResponse(w, json)
}

func (server *EventServer) startEventSender() {
        fmt.Println("Starting event sender goroutine...")
        for {
                data := <-server.eventQueue
                log.Println("About to send event: " + string(data))
                server.broadcastDEvent(data)
                time.Sleep(time.Millisecond * 50)
        }
}

func (server *EventServer) broadcastDEvent(data []byte) {
        //for index, wsConn := range server.connectionRegistry {
        for index := len(server.connectionRegistry) - 1; index > -1 ; index-- {
                wsConn := server.connectionRegistry[index]
                err := wsConn.WriteMessage(1, data)
                if err != nil {
                        // Detected disconnected channel. Need to clean up.
                        fmt.Printf("Could not write to channel: %v", err)
                        wsConn.Close()
                        server.removeConnection(index)
                }
        }
}

func (server *EventServer) removeConnection(i int)  {
        fmt.Printf("Removing index %v\n", i)

        copy(server.connectionRegistry[i:], server.connectionRegistry[i+1:])
        server.connectionRegistry[len(server.connectionRegistry)-1] = nil // or the zero value of T
        server.connectionRegistry = server.connectionRegistry[:len(server.connectionRegistry)-1]

        // server.connectionRegistry[len(server.connectionRegistry)-1], server.connectionRegistry[i] = server.connectionRegistry[i], server.connectionRegistry[len(server.connectionRegistry)-1]
        //server.connectionRegistry = server.connectionRegistry[:len(server.connectionRegistry)-1]
}

func (server *EventServer) registerChannel(w http.ResponseWriter, r *http.Request) {
        if r.URL.Path != "/start" {
                http.Error(w, "Not found", 404)
                return
        }
        if r.Method != "GET" {
                http.Error(w, "Method not allowed", 405)
                return
        }
        header := make(map[string][]string)

        header["Access-Control-Allow-Origin"] = []string{"*"}
        c, err := server.upgrader.Upgrade(w, r, header)
        if err != nil {
                log.Print("upgrade:", err)
                return
        }
        server.addConnection(c)
}

func (server *EventServer) addConnection(conn *websocket.Conn) {
        server.connectionRegistry = append(server.connectionRegistry, conn)
}

func writeResponse(w http.ResponseWriter, json []byte) {
        w.Header().Set("Content-Type", "application/json")
        w.Header().Set("Content-Length", strconv.Itoa(len(json)))
        w.Header().Set("Access-Control-Allow-Origin", "*")
        w.WriteHeader(http.StatusOK)
        w.Write(json)
}

// Handles Ctrl+C or most other means of "controlled" shutdown gracefully. Invokes the supplied func before exiting.
func handleSigterm(handleExit func()) {
        c := make(chan os.Signal, 1)
        signal.Notify(c, os.Interrupt)
        signal.Notify(c, syscall.SIGTERM)
        go func() {
                <-c
                handleExit()
                os.Exit(1)
        }()
}


