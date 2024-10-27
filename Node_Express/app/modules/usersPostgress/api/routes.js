import Controller from './controller.js';
import CustomRouter from '../../../pkg/custom/routes/Router.js';

const controller = new Controller()

// http://localhost:8046/v2/users/

class UserRouterPostgresSQL extends CustomRouter {
  constructor() {
    super();
    this.addRoute('get', '/', [], controller.get);
    this.addRoute('get', '/:id', [], controller.getById);
    this.addRoute('post', '/', [], controller.create);
    this.addRoute('put', '/:id', [], controller.updateId);
    this.addRoute('delete', '/:id', [], controller.deleteId);
  }
}

export default new UserRouterPostgresSQL().getRouter();