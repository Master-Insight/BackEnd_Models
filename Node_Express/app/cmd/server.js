import express from 'express';
import path from 'node:path'
import configEnv from '../pkg/services/env/env.js';
import cors from 'cors'
import __dirname from '../pkg/utilities/dirname.js';
import initializeDatabases from '../pkg/services/db/init.js';
import initializeMiddlewares from '../pkg/middleware/init.js';
import { logger } from '../pkg/middleware/logger.js';
import initializePassport from '../modules/auth/config/passport.config.js';
import passport from 'passport';
import appRouter from '../modules/routes.js'
import handleErrors from '../pkg/middleware/handleErrors.js';
import dotenv from 'dotenv';

// * App initialization ------------------------------
const app = express();

// * App Configurations --------------------------------
app.use(cors({origin:configEnv.cors_origin}));
app.use(express.json());
app.use(express.urlencoded({ extended: true }));
app.use(express.static(path.join(__dirname, 'public')))

// usado por deployer
dotenv.config()
const port = process.env.PORT || configEnv.config.port || 8080;

// App Data Source Configuration --------------------------------
const persistence = configEnv.services.persistence.service || "mongo"
initializeDatabases(persistence)

// App Middleware --------------------------------
initializeMiddlewares(app)

// passport
initializePassport()
app.use(passport.initialize())

// App Routes --------------------------------
app.use('/', appRouter);

// Error Handling Middleware --------------------------------
app.use(handleErrors)

// App Launch --------------------------------
app.listen(port, () => { logger.info(`SERVER: running on port: ${port}`); });