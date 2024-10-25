import { connect } from "mongoose"
import { logger } from "../../../middleware/logger.js";
import configEnv from "../env/env.js";

class MongoSingleton {
  static instance
  constructor() {
    const uri = configEnv.services.persistence.mongo_uri
    connect(uri);
  }

  static getInstance() {
    if(!this.instance){
      logger.info('BD Connected to Mongo');
      return this.instance = new MongoSingleton();
    }
    logger.info('BD was already connected to Mongo');
    return this.instance;
  }
}

export const connectDb = async () => {
  try {
    MongoSingleton.getInstance()
  } catch(err) {
    logger.error(`MONGO CONNECT: ${err}`)
  }
}