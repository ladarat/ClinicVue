import { ActionTree, MutationTree, GetterTree, Module } from 'vuex';
import LoginResponse from '@/common/api/modules/user/models/LoginRersponse';
import LoginRequest from '@/common/api/modules/user/models/LoginRequest';
import UserService from '@/common/api/modules/user/user.service';
import { LOGIN } from './action.type';
import { SET_AUTH, SET_ERROR, PURGE_AUTH } from './mutations.type';
import { RootState } from '../..';
import { IS_AUTH } from './getters.type';

export interface AuthState {
    errors: any,
    loginResponse: LoginResponse,
    isAuthenticated: boolean
}

const state: AuthState = {
    errors: null,
    loginResponse: <LoginResponse>{},
    isAuthenticated: false
};

const actions: ActionTree<AuthState, RootState> = {
    [LOGIN](context: any, request: LoginRequest) {
        return new Promise(resolve => {
            UserService.login(request)
                .then(({ data }) => {
                    context.commit(SET_AUTH, data);
                    resolve(data);
                })
                .catch(({ response }) => {
                    context.commit(SET_ERROR, response.data.errors);
                });
        });
    }
};

const mutations: MutationTree<AuthState> = {
    [SET_ERROR](state: any, error: Error) {
        state.isAuthenticated = false
        state.loginResponse = {}
        state.errors = error
    },
    [SET_AUTH](state: any, loginResponse: LoginResponse) {
        state.isAuthenticated = true
        state.loginResponse = loginResponse
        state.errors = {}
    },
    [PURGE_AUTH](state: any) {
        state.isAuthenticated = false
        state.loginResponse = {}
        state.errors = {}
    }
};

const getters: GetterTree<AuthState, RootState> = {
    [IS_AUTH](state: any): boolean {
      return state.isAuthenticated
    }
};

const authStoreModule: Module<AuthState, RootState> = {
    state,
    actions,
    mutations,
    getters
};

export default authStoreModule
