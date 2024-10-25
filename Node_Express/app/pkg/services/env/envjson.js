import __dirname from '../../../libraries/utils/dirname.js';
import fs from 'node:fs';
import path from 'node:path';
import { logger } from '../../../middleware/logger.js';

// Carga de archivo Env
const LoadEnvJson = () => {
  try {
    // Lee el archivo config.json
    const configPath = path.join(__dirname, 'env.json');
    const rawConfig = fs.readFileSync(configPath); 
    const config = JSON.parse(rawConfig);
    logger.info('ENV: Cargado');

    // Retorna el contenido del archivo
    return config;
  } catch (error) {
    logger.error('Error al cargar env.json:', error);
    return {};  // Retorna un objeto vacío si ocurre un error
  }
};

export default LoadEnvJson
