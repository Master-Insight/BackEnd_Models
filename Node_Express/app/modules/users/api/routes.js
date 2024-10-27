// import { Router } from "express";
// import Controller from "./controller.js";
// import { clients, handleAuth, users } from "../../../pkg/middleware/handlePolicies.js";
// import { uploader } from "../../../pkg/middleware/multer.middleware.js";
// import { celebrate } from "celebrate";

// const router = Router();
// const controller = new Controller()

// // http://localhost:8080/v1/users/

// router
// // public
// .get    ('/associates', controller.getAssociates)
// .get    ('/associate/:username', controller.getAssociate)

// // user
// .get    ('/current',        handleAuth(users),   controller.getUserSession)
// .put    ('/current/update', handleAuth(users),   controller.currentUpdate)
// .put    ('/current/uploadphoto',  
//   handleAuth(users), 
//   uploader(5, ['image/jpeg', 'image/jpg', 'image/png'], true).single('photo'),
//   controller.uploadPhoto)
  
// // Admins
// .get    ('/',               handleAuth(['ADMIN']), controller.get)

// export default router

import Controller from "./controller.js";
import CustomRouter from '../../../pkg/custom/routes/Router.js';

const controller = new Controller()

// http://localhost:8046/v1/users/

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