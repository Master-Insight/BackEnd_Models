import DaoMongo from "../../../pkg/custom/dao/dao.mongoose.js";
import dataModel from "./model.js";

export default class ThisDaoMongo extends DaoMongo {
  constructor() {
    super(dataModel);
  }

  updateConection = async (filter) => await this.model.findOneAndUpdate(
    filter,
    { connection: Date.now() },
    { new: true, timestamps: false }
  )
}

// caso mongo
// getUser = async (filter) => {
//   return await this.dao.get(filter, { password: 0 });
// };

// caso postgress
// async getUser(filter) {
//   const user = await this.dao.getBy(filter);
//   if (user) {
//     delete user.password; // Excluir el password manualmente
//   }
//   return user;
// }