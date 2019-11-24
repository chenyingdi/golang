package basic

type Url struct {
	Username string
	Password string
	Host     string
	Port     string
	VHost    string
}

/*
	create a new Url
 */
func NewUrl(username, password, host, port, vhost string) *Url {
	return &Url{
		Username: username,
		Password: password,
		Host:     host,
		Port:     port,
		VHost:    vhost,
	}
}

func (u *Url) ParseUrl() string {
	return "amqp://" +
		u.Username + ":" +
		u.Password + "@" +
		u.Host + ":" +
		u.Port + "/" +
		u.VHost
}