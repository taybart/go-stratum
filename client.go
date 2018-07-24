package main

/*
SPEC: https://en.bitcoin.it/wiki/Stratum_mining_protocol
Methods (client to server)
mining.authorize("username", "password")
mining.capabilities({"notify":[], "set_difficulty":{}, "set_goal":{}, "suggested_target": "hex target"})
mining.extranonce.subscribe()
mining.get_transactions("job id")
mining.submit("username", "job id", "ExtraNonce2", "nTime", "nOnce")
mining.subscribe("user agent/version", "extranonce1")
mining.suggest_difficulty(preferred share difficulty Number)
mining.suggest_target("full hex share target")
...
Methods (server to client)
client.get_version()
client.reconnect("hostname", port, waittime)
client.show_message("human-readable message")
mining.notify(...)
mining.set_difficulty(difficulty)
mining.set_extranonce("extranonce1", extranonce2_size)
mining.set_goal("goal name", {"malgo": "SHA256d", ...})
*/

type Client struct{}

func (c *Client) GetVersion(args *string, reply *string) error {
	return nil
}

func (c *Client) Reconnect(args *string, reply *string) error {
	return nil
}

func (c *Client) ShowMessage(args *string, reply *string) error {
	return nil
}
