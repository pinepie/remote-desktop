package rodeo

import (
	"github.com/angrypie/remote-desktop/server/rodeo/wserver"
	"log"
)

/*

Client - is client-side application
Host - it is device to which Client are connectiong for remote accesss

*/

//Host want to register on server. data contains host login
//If host successfully registered, than send to client message "RIGESTER_SUCCESS".
//If host with specified login already exist on server, than send "REGISTER_FAIL"
//with "login exist" in data field
func actHostRegister(data *interface{}, client *wserver.Client) {
	log.Println("REGISTER STRARt")
	hostLogin := (*data).(string)
	new_host := NewHost(hostLogin, client)
	ok := hostConnections.Add(new_host)
	if !ok {
		client.SendJson(&Action{"REGISTER_FAIL", "login exist"})
	} else {
		client.SendJson(&Action{"REGISTER_SUCCESS", ""})
	}
	log.Println("REGISTER STRARt")
}

//Client requests data about avaliable host.
//Server sends message "AVALIABLE_HOSTS" with hostInfo in data field
func actGetHosts(data *interface{}, client *wserver.Client) {
	hosts := hostConnections.getInfo()
	client.SendJson(&Action{"AVALIABLE_HOSTS", hosts})
}

//Client connects to host. Client and host send messages to each other
//and server doesn't process these messages, just forwards them
func actSelectHost(data *interface{}, client *wserver.Client) {
	name, _ := (*data).(string)
	host, ok := hostConnections.GetByLogin(name)
	if !ok {
		client.SendJson(&Action{"SELECT_FAIL", "host not exist"})
		return
	}
	log.Println("--------------LOCK")
	host.Lock()
	log.Println("--------------LOCKED")
	defer host.UnLock()

	if host.Active {
		client.SendJson(&Action{"SELECT_FAIL", "host busy"})
		return
	}
	host.Active = true

	host.Conn.SendJson(&Action{"CLIENT_CONNECT", ""})
	host.Wait()
	if host.Active == false {
		client.SendJson(&Action{"SELECT_FAIL", "denied"})
		return
	}

	host.Conn.SetOnmessage(copyMessage(client))
	client.SetOnmessage(copyMessage(host.Conn))

	clientConnected[client] = host
	client.SendJson(&Action{"SELECT_SUCCESS", ""})
}

func actClientAccess(data *interface{}, client *wserver.Client) {
	host := hostConnections.GetByConn(client)
	host.Active = true
	host.Signal()
}

func actClientDenied(data *interface{}, client *wserver.Client) {
	host := hostConnections.GetByConn(client)
	host.Active = false
	host.Signal()
}
