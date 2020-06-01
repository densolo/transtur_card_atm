

export function listenMessages(store: any) {
    let astilectron = window['astilectron'];
    astilectron.onMessage((message) => {
        console.log("Message: " + message.name + " payload: " + JSON.stringify(message.payload));

        if (message.name == 'update') {
            let cardState = JSON.parse(message.payload);

            store.dispatch({ 
                type: 'UPDATE',
                cardStateColor: cardState.card_state_color,
                cardStateText: cardState.card_state_text
            });
        }
    });  
}



