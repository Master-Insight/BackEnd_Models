import ThisDaoMongo from "../data/dao.mongo.js";
import { MongoService } from "../../../pkg/custom/service/service.mongoose.js";


export default class Service extends MongoService {
  constructor() {
    super(new ThisDaoMongo);
  }

  get = async (filter = {}, projection = null, options = {}) => await this.dao.get(filter, { password: 0, ...projection}, options);
  getBy = async (filter, projection = null) => await this.dao.getBy(filter, { password: 0, ...projection});

  updateConection = async (filter)  => await this.dao.updateConection(filter)

  // ACTUALIZACION DE IMAGEN
  updatePhoto = async (uid, path) => await this.dao.update({_id: uid}, {photo: path})
}