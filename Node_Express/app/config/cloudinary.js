import { v2 as cloudinary } from 'cloudinary'
import configEnv from '../services/env/env.js'

cloudinary.config({
  cloud_name: configEnv.services.images.cloudinary.name,
  api_key: configEnv.services.images.cloudinary.key,
  api_secret: configEnv.services.images.cloudinary.secret,
})

export default cloudinary