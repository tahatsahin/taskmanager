package common

func StartUp() {
	// initialize AppConfig variable
	initConfig()
	// initialize private/public keys for JWT auth
	// initKeys()
	// start a mongodb session
	createDbSession()
	// add indexes into mongodb
	addIndexes()
}
