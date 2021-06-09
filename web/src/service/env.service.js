function getEnvVal(name) {
        return process.env[`REACT_APP_${name}`] || '';
}