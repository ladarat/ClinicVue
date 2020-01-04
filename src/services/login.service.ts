import axios from 'axios';
import qs from 'qs';
import { RequestLogin } from '../models/requestLogin';

const login = (loginRes: RequestLogin) => axios({
  method: 'POST', url: 'http://192.168.1.62:8888/login', data: qs.stringify(loginRes),
}).then(() => {
  console.log('success');
});

export default {
  login,
};
