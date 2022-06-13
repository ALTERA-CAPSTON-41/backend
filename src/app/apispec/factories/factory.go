package apispec_factories

import apispec_handlers "clinic-api/src/app/apispec/handlers"

func Factory() apispec_handlers.Handler {
	return *apispec_handlers.NewHandler()
}
