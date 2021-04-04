package main

func main() {
	initLoggers()
	InfoLogger.Println("This in an information")
	WarningLogger.Println("Multiple modules are not imported")
	ErrorLogger.Println("Database handler failed")
}