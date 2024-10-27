// models/userModel.js

import { DataTypes } from 'sequelize';
import PostgresSingleton from '../../../pkg/services/db/connectPostgresSQL.js';
import { DOCTYPE, ROLES } from '../../utils/valueList.js';

// Obtén la conexión existente
const sequelize = PostgresSingleton.getInstance().sequelize;

const User = sequelize.define('User', {
  given_name: { type: DataTypes.STRING, allowNull: false, maxLength: 50 },
  family_name: { type: DataTypes.STRING, allowNull: false, maxLength: 50 },
  full_name: { type: DataTypes.STRING },
  username: { type: DataTypes.STRING, unique: true },
  email: { 
    type: DataTypes.STRING,
    allowNull: false,
    unique: true,
    validate: { isEmail: true },
  },
  password: { type: DataTypes.STRING, allowNull: false },
  role: { 
    type: DataTypes.ENUM(...ROLES),
    defaultValue: 'User',
  },
  document: { type: DataTypes.STRING, maxLength: 15 },
  documenttype: { type: DataTypes.ENUM(...DOCTYPE) },

  phone: { type: DataTypes.STRING, maxLength: 20 },
  photo: { type: DataTypes.STRING },
  bio: { type: DataTypes.STRING },
  birthday: { type: DataTypes.DATE },
  public: { type: DataTypes.BOOLEAN, defaultValue: true },

  linkedinId: { type: DataTypes.STRING },
  linkedinVerified: { type: DataTypes.BOOLEAN },

  created: { type: DataTypes.DATE, defaultValue: DataTypes.NOW },
  updated: { type: DataTypes.DATE, defaultValue: DataTypes.NOW },
  connection: { type: DataTypes.DATE, defaultValue: DataTypes.NOW },
}, {
  timestamps: true,
  createdAt: 'created',
  updatedAt: 'updated',
});

export default User;
