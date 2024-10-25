import { addLogger } from "./logger.js";
import handleResponses from "./handleResponses.js";

const initializeMiddlewares = async (app) => {
  app.use(addLogger)
  app.use(handleResponses)
};

export default initializeMiddlewares;