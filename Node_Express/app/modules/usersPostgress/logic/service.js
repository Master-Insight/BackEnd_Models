import { PostgresService } from "../../../pkg/custom/service/service.postgress.js";
import UserDao from "../data/dao.js";

class UserService extends PostgresService {
  constructor() {
    super(new UserDao());
  }
}

export default UserService;