import { Router } from "express";
import usersRouter from "./users/api/routes.js";
import usersSQLRouter from "./usersPostgress/api/routes.js";
import authRouter from "./auth/api/routes.js";
import valuesRouter from "./utils/routes.js";
import AppError from "../pkg/errors/AppError.js";

const router = Router()

// http://localhost:8080/

router.get('/', (req, res) => { res.send({"message": "Hello, Node.Js Express Backend!"}) });
router.use('/v1/users/', usersRouter)
router.use('/v1/auth/', authRouter)
router.use('/v1/values/', valuesRouter)
router.get('/v1/pruebas', async (req, res, next) => {res.send("Prueba Pruebas")});
router.all('*', (req, res, next) => { next(new AppError(`No se encuentra la url: ${req.originalUrl} en este servidor`, 404)); });

router.use('/v2/users/', usersSQLRouter)
export default router