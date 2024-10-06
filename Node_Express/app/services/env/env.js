import { logger } from "../../middleware/logger.js";
import LoadEnv from "./envfile.js";
import LoadEnvJson from "./envjson.js";

let configEnv = {};

// Load enviroment variables (method available "env" or "json")
export const LoadEnviroment = (method) => {
  switch (method) {
    case "env":
      Object.assign(configEnv, LoadEnv());
      break;
    
    case "json":
      Object.assign(configEnv, LoadEnvJson());
      break;

    default:
      logger.error("Unsupported environment method: " + method)
      break;
  }
}

export default configEnv;
