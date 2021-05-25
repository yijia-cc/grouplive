const Payment = () => {
  return (
    <div class="container">
          <main>
            <div class="py-5 text-center">
              <h2>Make a payment</h2>
              <p class="lead">The processing of payment will take time. You will receive an email upon the changes of status.</p>
            </div>
        
            <div class="row g-5">
              <div class="col-md-5 col-lg-4 order-md-last">
                <h4 class="d-flex justify-content-between align-items-center mb-3">
                  <span class="text-primary">Balance due</span>
                </h4>
                <ul class="list-group mb-3">
                  
                  <li class="list-group-item d-flex justify-content-between lh-sm">
                    <div>
                      <h6 class="my-0">Montly-Rent</h6>
                      <small class="text-muted"></small>
                    </div>
                    <span class="text-muted">$5</span>
                  </li>

                  <li class="list-group-item d-flex justify-content-between lh-sm">
                    <div>
                      <h6 class="my-0">Utilities</h6>
                      <small class="text-muted">Including water, gas, electricity</small>
                    </div>
                    <span class="text-muted">$5</span>
                  </li>

                  <li class="list-group-item d-flex justify-content-between lh-sm">
                    <div>
                      <h6 class="my-0">Previous Unpaid Rent</h6>
                      <small class="text-muted">Including interest fees</small>
                    </div>
                    <span class="text-muted">$10</span>
                  </li>
                  
                  <li class="list-group-item d-flex justify-content-between">
                    <span>Total (USD)</span>
                    <strong>$20</strong>
                  </li>
                </ul>
        

              </div>
              <div class="col-md-7 col-lg-8">
                <h4 class="mb-3">Payment Info</h4>
                <form class="needs-validation" novalidate>
                  <div class="row g-3">
                    <div class="col-sm-6">
                      <label for="firstName" class="form-label">First name</label>
                      <input type="text" class="form-control" id="firstName" placeholder="" value="" required></input>
                      <div class="invalid-feedback">
                        Valid first name is required.
                      </div>
                    </div>
        
                    <div class="col-sm-6">
                      <label for="lastName" class="form-label">Last name</label>
                      <input type="text" class="form-control" id="lastName" placeholder="" value="" required></input>
                      <div class="invalid-feedback">
                        Valid last name is required.
                      </div>
                    </div>
        
        
                    <div class="col-md-3">
                      <label for="payment" class="form-label">Amount paying</label>
                      <input type="number" min="0.00" step="0.01" max="25000" required></input>
                      <div class="invalid-feedback">
                        Please enter a valid amount of payment
                      </div>
                    </div>

                    

    
                  </div>
                  <hr></hr>
                  </form>
                  <button class="w-100 btn btn-primary btn-lg" type="submit">Pay with Stripe now!</button>

              </div>
            </div>
          </main>
        </div>
  );
};

export default Payment;
