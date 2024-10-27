import ThisDaoMongo from "../data/dao.mongo.js";
import { MongoService } from "../../../pkg/custom/service/service.mongoose.js";


export default class Service extends MongoService {
  constructor() {
    super(new ThisDaoMongo);
  }

  updateConection = async (filter)  => await this.dao.updateConection(filter)

  // ACTUALIZACION DE IMAGEN
  updatePhoto = async (uid, path) => await this.dao.update({_id: uid}, {photo: path})
}