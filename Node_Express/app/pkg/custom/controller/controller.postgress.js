/*
 * PostgresController: agrega métodos avanzados de PostgreSQL

 * CRUD
    get
    getBy + getById
    create 
    updateId (update)
    deleteId (delete)
 * AUXLIARES
    exists
    getUniqueValues
    getPaginate
    count
 * METODOS EN LOTE
    insertMany
    updateMany
    deleteMany
 * AVANZADAS
    query (solo postgres)
 */

import CustomController from "./controller.js";

export class PostgresController extends CustomController {
  constructor(service, fieldAllowed = []) {
    super(service, fieldAllowed);
  }
  
  // Método de consulta avanzada específico de PostgreSQL
  query = async (req, res, next) => {
    try {
      const { sqlQuery, replacements } = req.body;
      const result = await this.service.query(sqlQuery, replacements);
      res.sendSuccess(result);
    } catch (error) {
      next(error);
    }
  };
}