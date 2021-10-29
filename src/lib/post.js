import axios from 'axios'

export default function post(link, param) {
  let prefix = "http://localhost:8000/backend"
  return axios.post(prefix + link, param)
}