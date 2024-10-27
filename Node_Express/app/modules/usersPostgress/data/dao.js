import DaoPostgres from "../../../pkg/custom/dao/dao.postgress.js";
import User from "./model.js";

class UserDao extends DaoPostgres {
  constructor() {
    super(User);
  }
}

export default UserDao;