import 'package:flutter/material.dart';
import '../services/cart_service.dart';
import '../services/order_service.dart';
import '../models/cart.dart';

class CheckoutScreen extends StatefulWidget {
  const CheckoutScreen({super.key});

  @override
  State<CheckoutScreen> createState() => _CheckoutScreenState();
}

class _CheckoutScreenState extends State<CheckoutScreen> {
  final _formKey = GlobalKey<FormState>();
  final _addressController = TextEditingController();
  final _notesController = TextEditingController();
  final CartService _cartService = CartService();
  final OrderService _orderService = OrderService();

  List<Cart> _cartItems = [];
  bool _isLoading = true;
  bool _isProcessing = false;
  String _selectedPayment = 'bank_transfer';

  final List<Map<String, dynamic>> _paymentMethods = [
    {
      'value': 'bank_transfer',
      'label': 'Bank Transfer',
      'icon': Icons.account_balance,
    },
    {
      'value': 'e_wallet',
      'label': 'E-Wallet',
      'icon': Icons.account_balance_wallet,
    },
    {
      'value': 'cod',
      'label': 'Cash on Delivery',
      'icon': Icons.money,
    },
  ];

  @override
  void initState() {
    super.initState();
    _loadCart();
  }

  @override
  void dispose() {
    _addressController.dispose();
    _notesController.dispose();
    super.dispose();
  }

  Future<void> _loadCart() async {
    setState(() => _isLoading = true);
    try {
      final result = await _cartService.getCart();
      if (result['success'] && mounted) {
        setState(() {
          _cartItems = result['carts'] ?? [];
          _isLoading = false;
        });
      }
    } catch (e) {
      if (mounted) {
        setState(() => _isLoading = false);
      }
    }
  }

  double get _totalAmount {
    return _cartItems.fold(0, (sum, item) => sum + item.subtotal);
  }

  String get _formattedTotal {
    return 'Rp ${_totalAmount.toStringAsFixed(0).replaceAllMapped(
          RegExp(r'(\d{1,3})(?=(\d{3})+(?!\d))'),
          (Match m) => '${m[1]}.',
        )}';
  }

  Future<void> _placeOrder() async {
    if (!_formKey.currentState!.validate()) return;

    setState(() => _isProcessing = true);

    try {
      final result = await _orderService.createOrder(
        paymentMethod: _selectedPayment,
        shippingAddress: _addressController.text.trim(),
        notes: _notesController.text.trim(),
      );

      if (result['success'] && mounted) {
        ScaffoldMessenger.of(context).showSnackBar(
          const SnackBar(
            content: Text('Order placed successfully!'),
            backgroundColor: Colors.green,
          ),
        );
        Navigator.pushNamedAndRemoveUntil(context, '/home', (route) => false);
      } else {
        throw Exception(result['message']);
      }
    } catch (e) {
      if (mounted) {
        ScaffoldMessenger.of(context).showSnackBar(
          SnackBar(
            content: Text('Error: $e'),
            backgroundColor: Colors.red,
          ),
        );
      }
    } finally {
      if (mounted) {
        setState(() => _isProcessing = false);
      }
    }
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      backgroundColor: const Color(0xFF1A1A2E),
      appBar: AppBar(
        backgroundColor: const Color(0xFF2A2A4E),
        elevation: 0,
        leading: IconButton(
          icon: const Icon(Icons.arrow_back, color: Colors.white),
          onPressed: () => Navigator.pop(context),
        ),
        title: const Text(
          'Checkout',
          style: TextStyle(color: Colors.white, fontWeight: FontWeight.bold),
        ),
      ),
      body: _isLoading
          ? const Center(
              child: CircularProgressIndicator(
                valueColor: AlwaysStoppedAnimation<Color>(Color(0xFF6B5FED)),
              ),
            )
          : Column(
              children: [
                Expanded(
                  child: SingleChildScrollView(
                    padding: const EdgeInsets.all(20),
                    child: Form(
                      key: _formKey,
                      child: Column(
                        crossAxisAlignment: CrossAxisAlignment.start,
                        children: [
                          _buildSectionTitle('Order Summary'),
                          const SizedBox(height: 12),
                          _buildOrderSummary(),
                          const SizedBox(height: 24),
                          _buildSectionTitle('Payment Method'),
                          const SizedBox(height: 12),
                          _buildPaymentMethods(),
                          const SizedBox(height: 24),
                          _buildSectionTitle('Shipping Address'),
                          const SizedBox(height: 12),
                          _buildAddressField(),
                          const SizedBox(height: 24),
                          _buildSectionTitle('Order Notes (Optional)'),
                          const SizedBox(height: 12),
                          _buildNotesField(),
                        ],
                      ),
                    ),
                  ),
                ),
                _buildBottomBar(),
              ],
            ),
    );
  }

  Widget _buildSectionTitle(String title) {
    return Text(
      title,
      style: const TextStyle(
        color: Colors.white,
        fontSize: 18,
        fontWeight: FontWeight.bold,
      ),
    );
  }

  Widget _buildOrderSummary() {
    return Container(
      padding: const EdgeInsets.all(16),
      decoration: BoxDecoration(
        color: const Color(0xFF2A2A4E),
        borderRadius: BorderRadius.circular(16),
      ),
      child: Column(
        children: [
          ...(_cartItems.map((cart) => Padding(
                padding: const EdgeInsets.only(bottom: 12),
                child: Row(
                  children: [
                    Expanded(
                      child: Text(
                        '${cart.product.name} x${cart.quantity}',
                        style: const TextStyle(color: Colors.white),
                      ),
                    ),
                    Text(
                      cart.formattedSubtotal,
                      style: const TextStyle(
                        color: Colors.white,
                        fontWeight: FontWeight.bold,
                      ),
                    ),
                  ],
                ),
              ))),
          const Divider(color: Colors.white24),
          Row(
            mainAxisAlignment: MainAxisAlignment.spaceBetween,
            children: [
              const Text(
                'Total',
                style: TextStyle(
                  color: Colors.white,
                  fontSize: 16,
                  fontWeight: FontWeight.bold,
                ),
              ),
              Text(
                _formattedTotal,
                style: const TextStyle(
                  color: Color(0xFF6B5FED),
                  fontSize: 18,
                  fontWeight: FontWeight.bold,
                ),
              ),
            ],
          ),
        ],
      ),
    );
  }

  Widget _buildPaymentMethods() {
    return Column(
      children: _paymentMethods.map((method) {
        final isSelected = _selectedPayment == method['value'];
        return GestureDetector(
          onTap: () {
            setState(() => _selectedPayment = method['value']);
          },
          child: Container(
            margin: const EdgeInsets.only(bottom: 12),
            padding: const EdgeInsets.all(16),
            decoration: BoxDecoration(
              color: const Color(0xFF2A2A4E),
              borderRadius: BorderRadius.circular(12),
              border: Border.all(
                color: isSelected
                    ? const Color(0xFF6B5FED)
                    : Colors.white.withOpacity(0.1),
                width: isSelected ? 2 : 1,
              ),
            ),
            child: Row(
              children: [
                Icon(
                  method['icon'],
                  color: isSelected ? const Color(0xFF6B5FED) : Colors.white54,
                ),
                const SizedBox(width: 12),
                Expanded(
                  child: Text(
                    method['label'],
                    style: TextStyle(
                      color: isSelected ? Colors.white : Colors.white70,
                      fontWeight:
                          isSelected ? FontWeight.bold : FontWeight.normal,
                    ),
                  ),
                ),
                if (isSelected)
                  const Icon(Icons.check_circle, color: Color(0xFF6B5FED)),
              ],
            ),
          ),
        );
      }).toList(),
    );
  }

  Widget _buildAddressField() {
    return TextFormField(
      controller: _addressController,
      maxLines: 3,
      style: const TextStyle(color: Colors.white),
      validator: (value) {
        if (value == null || value.trim().isEmpty) {
          return 'Shipping address is required';
        }
        return null;
      },
      decoration: InputDecoration(
        hintText: 'Enter your complete shipping address',
        hintStyle: TextStyle(color: Colors.white.withOpacity(0.5)),
        filled: true,
        fillColor: const Color(0xFF2A2A4E),
        border: OutlineInputBorder(
          borderRadius: BorderRadius.circular(12),
          borderSide: BorderSide.none,
        ),
        enabledBorder: OutlineInputBorder(
          borderRadius: BorderRadius.circular(12),
          borderSide: BorderSide(color: Colors.white.withOpacity(0.1)),
        ),
        focusedBorder: OutlineInputBorder(
          borderRadius: BorderRadius.circular(12),
          borderSide: const BorderSide(color: Color(0xFF6B5FED), width: 2),
        ),
        errorBorder: OutlineInputBorder(
          borderRadius: BorderRadius.circular(12),
          borderSide: const BorderSide(color: Colors.red),
        ),
      ),
    );
  }

  Widget _buildNotesField() {
    return TextFormField(
      controller: _notesController,
      maxLines: 2,
      style: const TextStyle(color: Colors.white),
      decoration: InputDecoration(
        hintText: 'Add notes for your order (optional)',
        hintStyle: TextStyle(color: Colors.white.withOpacity(0.5)),
        filled: true,
        fillColor: const Color(0xFF2A2A4E),
        border: OutlineInputBorder(
          borderRadius: BorderRadius.circular(12),
          borderSide: BorderSide.none,
        ),
        enabledBorder: OutlineInputBorder(
          borderRadius: BorderRadius.circular(12),
          borderSide: BorderSide(color: Colors.white.withOpacity(0.1)),
        ),
        focusedBorder: OutlineInputBorder(
          borderRadius: BorderRadius.circular(12),
          borderSide: const BorderSide(color: Color(0xFF6B5FED), width: 2),
        ),
      ),
    );
  }

  Widget _buildBottomBar() {
    return Container(
      padding: const EdgeInsets.all(20),
      decoration: BoxDecoration(
        color: const Color(0xFF2A2A4E),
        borderRadius: const BorderRadius.only(
          topLeft: Radius.circular(24),
          topRight: Radius.circular(24),
        ),
        boxShadow: [
          BoxShadow(
            color: Colors.black.withOpacity(0.3),
            blurRadius: 10,
            offset: const Offset(0, -5),
          ),
        ],
      ),
      child: SafeArea(
        child: SizedBox(
          width: double.infinity,
          height: 56,
          child: ElevatedButton(
            onPressed: _isProcessing ? null : _placeOrder,
            style: ElevatedButton.styleFrom(
              backgroundColor: const Color(0xFF6B5FED),
              shape: RoundedRectangleBorder(
                borderRadius: BorderRadius.circular(16),
              ),
              disabledBackgroundColor: Colors.grey,
            ),
            child: _isProcessing
                ? const SizedBox(
                    width: 24,
                    height: 24,
                    child: CircularProgressIndicator(
                      strokeWidth: 2,
                      valueColor: AlwaysStoppedAnimation<Color>(Colors.white),
                    ),
                  )
                : const Text(
                    'Place Order',
                    style: TextStyle(
                      fontSize: 16,
                      fontWeight: FontWeight.bold,
                      color: Colors.white,
                    ),
                  ),
          ),
        ),
      ),
    );
  }
}




