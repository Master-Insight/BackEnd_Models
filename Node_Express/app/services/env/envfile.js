import dotenv from 'dotenv';
dotenv.config()

// Carga de archivo Env
const LoadEnv = () => {
  return {
    // APP CONFIG
    config : {
      name:         process.env.APP_NAME,
      port:         process.env.PORT,
    },
    // CORS
    cors_origin:  process.env.CORS_ORIGIN,
    
    // ADMIN
    admin: {
      users:      process.env.USERS_ADMIN,
      pass:  process.env.USER_ADMIN_PASS,
    },

    codes: {
      jwt:     process.env.JWT_SECRET_CODE,
    },

    // SERVICES
    services: {
      persistence: {
        service: process.env.PERSISTENCE,
        mongo: [
          {
            uri:    process.env.MONGO_URI,
          }
        ],
      },
      email: {
        user_name:  process.env.USER_NAME,
        gmail_user:   process.env.GMAIL_USER,
        gmail_pass:   process.env.GMAIL_PASS,
      },
      images: {
        cloudinary: {
          name:    process.env.CLOUDINARY_NAME,
          key:     process.env.CLOUDINARY_KEY,
          secret:  process.env.CLOUDINARY_SECRET,
          url:     process.env.CLOUDINARY_URL,
          default_folder: process.env.CLOUDINARY_DEFAULT_FOLDER,
        }
      }
    }
  };
};

export default LoadEnv;