import { Router } from 'express';
import UserController from './controller.js';

const router = Router();
const userController = new UserController();

router.get('/', userController.get);
router.get('/:eid', userController.getById);
router.post('/', userController.create);
router.put('/:eid', userController.updateId);
router.delete('/:eid', userController.deleteId);

export default router;