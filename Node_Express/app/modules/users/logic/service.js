import CustomService from "../../../pkg/custom/service.js";
import ThisDaoMongo from "../data/dao.mongo.js";
import AppError from "../../../pkg/errors/AppError.js";


export default class Service extends CustomService {
  constructor() {
    super(new ThisDaoMongo);
  }

  get = async (filter, excludePassword = true )  => await this.dao.get   (filter, excludePassword)
  getBy = async (filter, excludePassword = true) => await this.dao.getBy (filter, excludePassword)

  // ACTUALIZACION DE IMAGEN
  updatePhoto = async (uid, path) => {
    return await this.dao.update({_id: uid}, {photo: path})
  }
}