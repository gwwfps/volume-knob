import board
from analogio import AnalogIn
import adafruit_dotstar as dotstar

dot = dotstar.DotStar(board.APA102_SCK, board.APA102_MOSI, 1, brightness=0.2)

analog2in = AnalogIn(board.D2)

def getPortion(pin):    
    return round(min(pin.value / 65400.0, 1), 2)

s = None
while True:
  p = getPortion(analog2in)  
  
  steps = int(round(p * 53))
  if not s is None and steps != s:
    print("Percent:%d" % int(p * 100))
    print("Step:%d" % steps)

  s = steps

  dot[0] = [0, 0, int(255 * p)]
