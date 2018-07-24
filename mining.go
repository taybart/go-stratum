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

type Mining struct{}

func (m *Mining) Authorize(args *string, reply *string) error {
	return nil
}

func (m *Mining) Capabilites(args *string, reply *string) error {
	return nil
}

func (m *Mining) GetTransactions(args *string, reply *string) error {
	return nil
}

func (m *Mining) Submit(args *string, reply *string) error {
	return nil
}

func (m *Mining) Subscribe(args *string, reply *string) error {
	return nil
}

func (m *Mining) SuggestDifficulty(args *string, reply *string) error {
	return nil
}

func (m *Mining) SuggestTarget(args *string, reply *string) error {
	return nil
}

func (m *Mining) Notify(args *string, reply *string) error {
	return nil
}

func (m *Mining) SetDifficulty(args *string, reply *string) error {
	return nil
}

func (m *Mining) SetExtranonce(args *string, reply *string) error {
	return nil
}

func (m *Mining) SetGoal(args *string, reply *string) error {
	return nil
}

// func (m *Mining) Extranonce(args *string, reply *string) error {
// }
