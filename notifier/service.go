package notifier

import (
	"fmt"
	"os/exec"
)

type Notifier struct{}

func NewNotifier() Notifier {
	return Notifier{}
}

func (n Notifier) Notify(notification, title, url string) error {
	msg := fmt.Sprintf(`display notification "%s" with title "%s"`, url, title)
	_, err := exec.Command("osascript", "-e", msg).Output()
	return err
}
