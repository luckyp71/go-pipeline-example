package example1

//Convert a list of emails (string) to a channel.
//The goroutine sends the emails on the channel and closes the channel when all values has been sent.
func Gen(emails ...string) <-chan string {
	out := make(chan string)
	go func() {
		for _, email := range emails {
			out <- email
		}
		close(out)
	}()
	return out
}

//A function that simulate sending email
func SendEmail(input <-chan string) <-chan string {
	out := make(chan string)
	go func() {
		for email := range input {
			out <- "Sending email to " + email
		}
		close(out)
	}()
	return out
}
