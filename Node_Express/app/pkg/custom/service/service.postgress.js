/*
 * Servicio específico para PostgreSQL

 * CRUD
    get
    getBy
    create
    update
    delete
 * AUXLIARES
    exists
    getUniquesValues
    getPaginate
    count
 * MÉTODOS EN LOTE
    insertMany 
    updateMany 
    deleteMany 
 * AVANZADAS
    aggregate (solo mongodb)
 */

import CustomService from './service.js';

export class PostgresService extends CustomService {
  constructor(dao) {
    super(dao);
  }

  // Método de consulta avanzada (solo PostgreSQL)
  query = async (sqlQuery, replacements = {}) => {
    return await this.dao.query(sqlQuery, replacements);
  };
}