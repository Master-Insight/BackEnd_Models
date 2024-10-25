import { Sequelize } from 'sequelize';
import { logger } from '../../middleware/logger.js'; // Ajusta la ruta según tu estructura
import configEnv from '../env/env.js'; // Ajusta la ruta según tu estructura

class PostgresSingleton {
  static instance;

  constructor() {
    // Inicializa Sequelize con la URI de conexión de Neon
    const uri = configEnv.services.persistence.sql_uri
    this.sequelize = new Sequelize(uri, {
      dialect: 'postgres',
      dialectOptions: {
        ssl: {
          require: true,
          rejectUnauthorized: false, // Esto es importante para Neon
        },
      },
    });
  }

  static getInstance() {
    if (!this.instance) {
      this.instance = new PostgresSingleton();
      logger.info('BD Connected to PostgreSQL');
    } else {
      logger.info('BD was already connected to PostgreSQL');
    }
    return this.instance;
  }

  async connect() {
    try {
      await this.sequelize.authenticate();
      logger.info('PostgreSQL connection is authenticated.');
    } catch (error) {
      logger.error('Unable to connect to the PostgreSQL database:', error);
    }
  }
}

export const connectDb = async () => {
  const postgresSingleton = PostgresSingleton.getInstance();
  await postgresSingleton.connect();
};

export default PostgresSingleton;
