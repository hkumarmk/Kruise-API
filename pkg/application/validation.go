package application

import (
	"bytes"
	"deploy-wizard/gen/models"
	"encoding/json"
	"fmt"
	"strconv"

	log "github.com/sirupsen/logrus"
)

const (
	errMsgInvalidJSON      = "invalid json payload"
	errMsgNotAnApplication = "not an application object"
)

// ValidateApplication returns of map with key = field and value = error
func ValidateApplication(appdata interface{}) map[string]interface{} {
	errors := map[string]interface{}{}

	var b bytes.Buffer
	enc := json.NewEncoder(&b)
	err := enc.Encode(appdata)
	if err != nil {
		log.WithField("f", "application.ValidateApplication").WithError(err).Warnf(errMsgInvalidJSON)
		errors[""] = errMsgInvalidJSON
		return errors

	}

	var app *models.Application
	err = json.Unmarshal(b.Bytes(), &app)
	if err != nil {
		log.WithField("f", "application.ValidateApplication").WithError(err).Warnf(errMsgNotAnApplication)
		errors[""] = errMsgNotAnApplication
		return errors
	}

	if app.Name == "" {
		errors["name"] = newRequiredValidationError("name")
	}

	if app.Release == "" {
		errors["release"] = newRequiredValidationError("release")
	}

	if app.Environment == "" {
		errors["environment"] = newRequiredValidationError("environment")
	}

	if app.Tenant == "" {
		errors["tenant"] = newRequiredValidationError("tenant")
	}

	if app.Namespace == "" {
		errors["namespace"] = newRequiredValidationError("namespace")
	}

	if app.Path == "" {
		errors["path"] = newRequiredValidationError("path")
	}

	if app.Region == "" {
		errors["region"] = newRequiredValidationError("region")
	}

	if app.RepoURL == "" {
		errors["repoURL"] = newRequiredValidationError("repoURL")
	}

	if app.TargetRevision == "" {
		errors["targetRevision"] = newRequiredValidationError("targetRevision")
	}

	if len(app.Services) > 0 {
		servicesErrors := ValidateServices(app.Services)
		if len(servicesErrors) > 0 {
			errors["services"] = servicesErrors
		}
	}

	return errors
}

// ValidateServices returns of map with key = field and value = error
func ValidateServices(services []*models.Service) map[string]interface{} {
	errors := map[string]interface{}{}

	for i, svc := range services {
		svcErrors := ValidateService(svc)
		idx := strconv.Itoa(i)

		if len(svcErrors) > 0 {
			errors[idx] = svcErrors
		}
	}

	return errors
}

// ValidateService returns of map with key = field and value = error
func ValidateService(svc *models.Service) map[string]interface{} {
	errors := map[string]interface{}{}

	if svc.Name == "" {
		errors["name"] = newRequiredValidationError("name")
	}

	if len(svc.Ports) == 0 {
		errors["ports"] = newRequiredValidationError("ports")
		return errors
	}

	portsErrors := ValidateServicePorts(svc.Ports)
	if len(portsErrors) > 0 {
		errors["ports"] = portsErrors
	}

	return errors
}

// ValidateServicePorts returns of map with key = field and value = error
func ValidateServicePorts(ports []*models.ServicePort) map[string]interface{} {
	errors := map[string]interface{}{}

	for i, port := range ports {
		portErrors := ValidateServicePort(port)
		idx := strconv.Itoa(i)

		if len(portErrors) > 0 {
			errors[idx] = portErrors
		}
	}

	return errors
}

// ValidateServicePort returns of map with key = field and value = error
func ValidateServicePort(port *models.ServicePort) map[string]interface{} {
	errors := map[string]interface{}{}

	if port.Name == "" {
		errors["name"] = newRequiredValidationError("name")
	}

	if port.Port == 0 {
		errors["port"] = newRequiredValidationError("port")
	}

	return errors
}

func newRequiredValidationError(field string) string {
	return fmt.Sprintf("%s is required", field)
}
