import { connect } from "mongoose"
import { logger } from "../middleware/logger.js";
import configEnv from "./env/env.js"


class MongoSingleton {
  static instance
  constructor() {
    connect(configEnv.services.prsistence.mongo_uri[0]);
  }

  static getInstance() {
    if(!this.instance){
      logger.info('BD Connected to Mongo');
      return this.instance = new MongoSingleton();
    }
    logger.info('BD was already connected');
    return this.instance;
  }
}

export const connectDb = async () => {
  try {
    MongoSingleton.getInstance()
  } catch(err) {
    logger.error(err)
  }
}