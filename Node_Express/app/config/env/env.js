import { logger } from "../../middleware/logger.js";
import LoadEnv from "./envfile.js";
import LoadEnvJson from "./envjson.js";

const configEnv = LoadEnvJson()

// hay un erro en esto por eso no lo uso
// Load enviroment variables (method available "env" or "json")
export const LoadEnviroment = (method) => {
  if (method == "env") {
    configEnv = LoadEnv()
  } else if (method == "json") {
    configEnv = LoadEnvJson()
  } else {
    logger.error("Unsupported environment method: " + method)
  }
}

export default configEnv;
