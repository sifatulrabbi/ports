import {configs} from "@/app/config"

export const getCoreApiUrl = (path = "", v = "v1") => {
	return configs.API_URL + "/api/" + v + path
}
