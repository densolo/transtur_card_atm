
import * as React from 'react';
import { ReactReduxContext } from 'react-redux'
import { connect, useSelector } from 'react-redux'
import {trEn, trRo} from 'translations/translation';


const AppForm = () => {
    const cardState = useSelector(state => state) as any;
  
    return (
        <div className="text-center" 
            style={{"position": "absolute", "top": "0", "bottom": "0", "left": "15px", "right": "15px"}}>
            <div className="row h-100">
                <div className="w-25"></div>
                <div className="w-50 align-self-center">
                    <div className={"alert alert-" + cardState.cardStateColor}>
                        {trEn(cardState.cardStateText)}
                    </div>
                    <div style={{"height": "100px"}}>

                    </div>
                    <div className={"alert alert-" + cardState.cardStateColor}>
                        {trRo(cardState.cardStateText)}
                    </div>
                </div>
                <div className="w-25"></div>
            </div>
        </div>
    );
  };
  
export default AppForm;
