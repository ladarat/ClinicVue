import ApiService from '@/common/api/api.service'
import LoginRequest from './models/LoginRequest'

const UserService = {
  login(loginRequest: LoginRequest) {
    return ApiService.post('', loginRequest)
  }
}

export default UserService