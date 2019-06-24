

const userReducer = (state = [], action) => {
    switch(action.type) {
        case 'ADD_USER':
          return state.concat([action.data]);
        case 'SHOW_USER':
          return 
        default:
          return state;
    }
}

export default userReducer;