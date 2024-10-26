/*
  * Este DAO es de PostgressSQL junto a Sequelize 
*/

export default class DaoPostgres {
  constructor(model) {
    this.model = model;
  }

  /*
   * Operaciones CRUD comunes: suelen ser los métodos más usados en cualquier DAO.
   */

  // Obtiene múltiples registros que coinciden con un filtro.
  get = async (filter = {}, options = {}) => {
    return await this.model.findAll({ where: filter, ...options });
  };

  // Obtiene un único registro que coincida con un filtro.
  getBy = async (filter) => {
    return await this.model.findOne({ where: filter });
  };

  // Crea un nuevo registro en la tabla.
  create = async (newElement) => {
    return await this.model.create(newElement);
  };

  // Actualiza un registro basado en un filtro y devuelve el registro actualizado.
  update = async (filter, elementUpdate) => {
    const [updatedCount, [updatedElement]] = await this.model.update(elementUpdate, {
      where: filter,
      returning: true,
    });
    return updatedCount > 0 ? updatedElement : null; // Retorna el elemento actualizado o null si no se encontró.
  };

  // Elimina un registro basado en un filtro.
  delete = async (filter) => {
    const elementToDelete = await this.getBy(filter);
    if (elementToDelete) {
      await this.model.destroy({ where: filter });
      return elementToDelete; // Retorna el registro eliminado.
    }
    return null; // Retorna null si no se encontró el registro.
  };

  /*
   * Métodos auxiliares: proporcionan funcionalidades adicionales comunes en proyectos.
   */

  // Verifica si existe un registro que coincida con un filtro.
  exists = async (filter) => {
    return !!(await this.getBy(filter));
  };

  // Obtiene valores únicos de un campo específico en la tabla.
  getUniquesValues = async (field) => {
    const results = await this.model.findAll({
      attributes: [[this.model.sequelize.fn('DISTINCT', this.model.sequelize.col(field)), field]],
    });
    return results.map((result) => result[field]); // Retorna los valores únicos como un array.
  };

  // Cuenta la cantidad de registros que coinciden con un filtro.
  count = async (filter = {}) => {
    return await this.model.count({ where: filter });
  };

  /*
   * Operaciones en lote: ayudan a manejar múltiples registros.
   */

  // Inserta múltiples registros de una sola vez.
  insertMany = async (elements) => {
    return await this.model.bulkCreate(elements);
  };

  // Actualiza múltiples registros que coincidan con un filtro.
  updateMany = async (filter, elementUpdate) => {
    const [updatedCount] = await this.model.update(elementUpdate, {
      where: filter,
    });
    return updatedCount; // Retorna el número de registros actualizados.
  };

  // Elimina múltiples registros que coincidan con un filtro.
  deleteMany = async (filter) => {
    return await this.model.destroy({ where: filter });
  };

  /*
   * Agregaciones: son menos frecuentes y suelen utilizarse en casos más avanzados.
   */

  // Realiza consultas avanzadas (a través de raw queries si es necesario).
  query = async (sqlQuery, replacements = {}) => {
    return await this.model.sequelize.query(sqlQuery, {
      replacements,
      type: this.model.sequelize.QueryTypes.SELECT,
    });
  };
}
