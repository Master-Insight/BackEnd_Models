/*
  * Este DAO es de MongoDB junto a Mongoose
*/

export default class DaoMongo {
  constructor(model) {
    this.model = model;
  }

  /*
   * Operaciones CRUD comunes: suelen ser los métodos más usados en cualquier DAO.
  */
  
  // Obtiene múltiples documentos que coinciden con un filtro. 
  get = async (filter = {}, projection = null, options = {}) => {
    return await this.model.find(filter, projection, options);
  };

  // Obtiene un único documento que coincida con un filtro.
  getBy = async (filter, projection = null) => {
    return await this.model.findOne(filter, projection);
  };

  // Crea un nuevo documento en la colección.
  create = async (newElement) => {
    return await this.model.create(newElement);
  };

  // Actualiza un documento basado en un filtro y devuelve el documento actualizado.
  update = async (filter, elementUpdate, options = { new: true }) => {
    return await this.model.findOneAndUpdate(filter, elementUpdate, options);
  };

  // Elimina un documento basado en un filtro.
  delete = async (filter, options = { new: true }) => {
    return await this.model.findOneAndDelete(filter, options);
  };

  /*
   * Métodos auxiliares: proporcionan funcionalidades adicionales comunes en proyectos.
  */

  // Verifica si existe un documento que coincida con un filtro.
  // Comúnmente usado en validaciones previas a la creación o actualización de datos.
  exists = async (filter) => {
    return !!(await this.getBy(filter));
  };

  // Obtiene valores únicos de un campo específico en la colección.
  // Útil para obtener listas de categorías, tags u otros valores únicos.
  getUniquesValues = async (field) => {
    return await this.model.distinct(field);
  };

  // Obtiene documentos paginados en base a un filtro.
  // Se usa para listas paginadas, en combinación con opciones de paginación.
  getPaginate = async (filter, options) => {
    return await this.model.paginate(filter, options);
  };

  // Cuenta la cantidad de documentos que coinciden con un filtro.
  // Útil para conocer el total de registros para paginación o estadísticas.
  count = async (filter = {}) => {
    return await this.model.countDocuments(filter);
  };

  /*
   * Operaciones en lote: ayudan a manejar múltiples documentos.
  */

  // Inserta múltiples documentos de una sola vez.
  // Ideal para cargas masivas o migraciones.
  insertMany = async (elements) => {
    return await this.model.insertMany(elements);
  };

  // Actualiza múltiples documentos que coincidan con un filtro.
  // Usado cuando varios registros necesitan un cambio similar.
  updateMany = async (filter, elementUpdate) => {
    return await this.model.updateMany(filter, elementUpdate, { new: true });
  };

  // Elimina múltiples documentos que coincidan con un filtro.
  // Ideal para limpiar datos en masa, como registros antiguos o duplicados.
  deleteMany = async (filter) => {
    return await this.model.deleteMany(filter);
  };

  /*
   * Agregaciones: son menos frecuentes y suelen utilizarse en casos más avanzados.
  */

  // Realiza una agregación avanzada con un pipeline de MongoDB.
  // Permite realizar operaciones complejas, como agrupaciones y filtros avanzados.
  aggregate = async (pipeline) => {
    return await this.model.aggregate(pipeline);
  };
}
