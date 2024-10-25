import { connectDb as connectMongo } from "./connectMongo.js";
import { connectDb as connectPostgres } from "./connectPostgresSQL.js";

const initializeDatabases = async () => {
    // Iniciar conexión a MongoDB
    await connectMongo();

    // Iniciar conexión a PostgreSQL
    await connectPostgres();
};

export default initializeDatabases;
