
let initialState = {
    cardStateColor: 'secondary',
    cardStateText: 'Unknown'
}


export default function stateUpdate(state, action: any) {

    if (typeof state === 'undefined') {
        return initialState;
    }
    
    console.log("stateUpdate: " + JSON.stringify(action));

    switch (action.type) {
      case 'UPDATE':
        return Object.assign({}, state, {
            cardStateColor: action.cardStateColor,
            cardStateText: action.cardStateText
        });

      default:
        return state;
    }
  }
  