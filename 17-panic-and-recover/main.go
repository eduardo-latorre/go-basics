package main

import "log"

// Panic and Recover
// - Everything under Panic won't be executed
// - When it's trigerred a Panic, it will look for any defer functions before the panic line
// - The defer function must always exist before the panic execution
// - The defere anonimous function will be executed if recover is not equal to nil which means whether panic has been executed or not

func accessingJediSecrets() string {
	defer func() {
		if r := recover(); r != nil {
			log.Println("closing every access to jedi secrets")
		}
	}()
	log.Println("accessing the jedi secrets")
	secret := gettingJediSecret()
	return secret
}

func gettingJediSecret() string {
	panic("we've detected a force disturbance")
	return "Might the force be with you"
}

func main() {
	if secret := accessingJediSecrets(); secret == "" {
		secret = "censored"
		log.Println("the secret is:", secret)
	}
	log.Println("end of the program")
}
