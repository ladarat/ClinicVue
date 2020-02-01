import ApiService from '@/common/api/api.service'
import LoginRequest from './models/LoginRequest'
import LoginResponse from './models/LoginRersponse'

const UserService = {
  login(loginRequest: LoginRequest) {
    return ApiService.post<LoginResponse>('', loginRequest)
  }
}

export default UserService