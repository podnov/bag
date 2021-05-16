import {
	GET_WALLET,
	GET_WALLET_ERROR,
	GET_WALLET_SUCCESS
} from '../types'

const initialState = {
	address: null,
	isLoading: false,
	statistics: null
}

export default function reduce(state = initialState, action){
	switch(action.type){
		case GET_WALLET:
			return {
				...state,
				isLoading: true
			};
		case GET_WALLET_SUCCESS:
			return {
				...state,
				statistics: action.payload,
				isLoading: false
			};
		default:
			return state;
	}

}
