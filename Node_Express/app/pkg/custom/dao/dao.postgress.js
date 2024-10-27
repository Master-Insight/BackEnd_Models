/*
 * DAO para PostgreSQL utilizando Sequelize

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
 * METODOS EN LOTE
    insertMany 
    updateMany 
    deleteMany 
 * AVANZADAS
    query  (solo postgress)

 */

export default class DaoPostgres {
  constructor(model) {
    this.model = model;
  }

  // Métodos CRUD
  get = async (filter = {}, options = {}) => {
    return await this.model.findAll({ where: filter, ...options });
  };

  getBy = async (filter) => {
    return await this.model.findOne({ where: filter });
  };

  create = async (newElement) => {
    return await this.model.create(newElement);
  };

  update = async (filter, elementUpdate) => {
    const [updatedCount, [updatedElement]] = await this.model.update(elementUpdate, {
      where: filter,
      returning: true,
    });
    return updatedCount > 0 ? updatedElement : null;
  };

  delete = async (filter) => {
    const elementToDelete = await this.getBy(filter);
    if (elementToDelete) {
      await this.model.destroy({ where: filter });
      return elementToDelete;
    }
    return null;
  };

  // Métodos auxiliares
  exists = async (filter) => {
    return !!(await this.getBy(filter));
  };

  getUniquesValues = async (field) => {
    const results = await this.model.findAll({
      attributes: [[this.model.sequelize.fn('DISTINCT', this.model.sequelize.col(field)), field]],
    });
    return results.map((result) => result[field]);
  };

  getPaginate = async (filter = {}, options = {}) => {
    const { limit = 10, offset = 0 } = options;
    return await this.model.findAndCountAll({
      where: filter,
      limit,
      offset,
    });
  };

  count = async (filter = {}) => {
    return await this.model.count({ where: filter });
  };

  // Métodos en lote
  insertMany = async (elements) => {
    return await this.model.bulkCreate(elements);
  };

  updateMany = async (filter, elementUpdate) => {
    const [updatedCount] = await this.model.update(elementUpdate, {
      where: filter,
    });
    return updatedCount;
  };

  deleteMany = async (filter) => {
    return await this.model.destroy({ where: filter });
  };

  /*
   * Agregaciones: son menos frecuentes y suelen utilizarse en casos más avanzados.
   */

  // Consultas avanzadas
  query = async (sqlQuery, replacements = {}) => {
    return await this.model.sequelize.query(sqlQuery, {
      replacements,
      type: this.model.sequelize.QueryTypes.SELECT,
    });
  };
}
