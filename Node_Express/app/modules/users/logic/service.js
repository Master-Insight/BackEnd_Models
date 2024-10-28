import ThisDaoMongo from "../data/dao.mongo.js";
import { MongoService } from "../../../pkg/custom/service/service.mongoose.js";


export default class Service extends MongoService {
  constructor() {
    super(new ThisDaoMongo);
  }

  // Función auxiliar para eliminar password de la proyección
  handlePasswordProjection = (projection) => {
    const finalProjection = {};
    
    if (projection) {
      if (Object.values(projection).some(value => value === 1)) {
        Object.assign(finalProjection, projection);
        delete finalProjection.password; // Eliminar password si se incluye un campo positivo
      } else {
        Object.assign(finalProjection, { password: 0, ...projection }); // Agregar password: 0
      }
    } else {
      finalProjection.password = 0; // Sin proyección, agregar password: 0
    }

    return finalProjection;
  };
  
  get = async (filter = {}, projection = null, options = {}) => {
    const finalProjection = this.handlePasswordProjection(projection);
    return await this.dao.get(filter, finalProjection, options);
  };
  getBy = async (filter, projection = null) => {
    const finalProjection = this.handleProjection(projection);
    return await this.dao.getBy(filter, finalProjection);
  };

  updateConection = async (filter)  => await this.dao.updateConection(filter)

  // ACTUALIZACION DE IMAGEN
  updatePhoto = async (uid, path) => await this.dao.update({_id: uid}, {photo: path})
}