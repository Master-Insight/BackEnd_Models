import CustomController from "../../../pkg/custom/controller/controller.js";
import UserService from "../logic/service.js";

class UserController extends CustomController {
  constructor() {
    const service = new UserService();
    super(service, ['username', 'email', 'role', 'documenttype']);
  }
}

export default UserController;