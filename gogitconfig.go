package gogitconfig

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os/exec"
	"strings"
	"syscall"
)

type ConfigValue struct {
	key string
}

func New(key string) *ConfigValue {
	return &ConfigValue{key}
}

func (val *ConfigValue) Get() (string, error) {
	return val.execCommand([]string{val.key})
}

func (val *ConfigValue) Set(configValue string) error {
	_, err := val.execCommand([]string{val.key, configValue})
	return err
}

func (val *ConfigValue) Unset() error {
	_, err := val.execCommand([]string{"--unset", val.key})
	return err
}

func (val *ConfigValue) UnsetGlobal() error {
	_, err := val.execCommand([]string{"--global", "--unset", val.key})
	return err
}

func (val *ConfigValue) SetGlobal(configValue string) error {
	_, err := val.execCommand([]string{"--global", val.key, configValue})
	return err
}

func (val *ConfigValue) execCommand(args []string) (string, error) {
	var stdout bytes.Buffer
	cmdArguments := []string{"config"}
	cmdArguments = append(cmdArguments, args...)
	cmd := exec.Command("git", cmdArguments...)
	cmd.Stdout = &stdout
	cmd.Stderr = ioutil.Discard
	err := cmd.Run()
	if exitError, ok := err.(*exec.ExitError); ok {
		if waitStatus, ok := exitError.Sys().(syscall.WaitStatus); ok {
			errorCode := waitStatus.ExitStatus()
			switch errorCode {
			case 0:
				{
					break
				}
			case 1:
				{
					return "", errors.New("gogitconfig: the section or key is invalid")
				}
			case 2:
				{
					return "", errors.New("gogitconfig: no section or name was provided")
				}
			case 3:
				{
					return "", errors.New("gogitconfig: the config file is invalid")
				}
			case 4:
				{
					return "", errors.New("gogitconfig: can not write to the config file")
				}
			case 5:
				{
					return "", errors.New("gogitconfig: you try to unset an option which does not exist, or you try to unset/set an option for which multiple lines match. Do you want to use GlobalUnset()?")
				}
			case 6:
				{
					return "", errors.New("gogitconfig: you try to use an invalid regexp")
				}
			default:
				{
					return "", errors.New("gogitconfig: git command returned unknown exit code (" + fmt.Sprintf("%v", errorCode) + ") ")
				}
			}
			if waitStatus.ExitStatus() == 1 {
				return "", errors.New("gogitconfig: key not found")
			} else if waitStatus.ExitStatus() == 5 {

			}
		}
		return "", err
	}
	return strings.TrimRight(stdout.String(), "\n\000"), nil
}

/*

func main() {
	configValue := New("test.value")
	v, err := configValue.Get()
	if err != nil {
		log.Println("error in Get:", err.Error())
		err := configValue.Set("foobar")
		if err != nil {
			log.Println("error in Set:", err.Error())
		}
	} else {
		log.Println("value: ", v)
		err = configValue.Unset()
		if err != nil {
			log.Println("error in Unset:", err.Error())
		}
	}
}

*/
