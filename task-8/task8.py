from PIL import Image #Pillow==6.1.0 
import pytesseract #pytesseract==0.2.7

expression = pytesseract.image_to_string(Image.open('eqn.png'), config='--psm 7')
print(expression) 
print(eval(expression))
