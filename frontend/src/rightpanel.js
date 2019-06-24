import React from 'react';

export class RightPanel extends React.Component {
    renderLogin(isLoggedIn) {
            if (isLoggedIn) {
                return
            }
            return (<div className="login">
                    <input placeholder="Email" type="email" onChange={this.props.onUsernameChange} />
                    <input placeholder="Password" value={this.props.password} type="password" onChange={this.props.onPasswdChange} />
                    <button onClick={this.props.onLogin}> Log in </button>
                    <br />
                    <a href="/aeilos/register.html">Click to register</a>
                </div>)
    }

    render() {
        let isLoggedIn = this.props.email !== '';
        
        return (
            <div>
                {this.renderLogin(isLoggedIn)}
                <div>
                    Logged in as: {this.props.email}
                    <button onClick={this.props.onLogOut}> Log out </button>
                    
                </div>
            </div>
        );
    }
}